package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	pb "chess_grpc/proto"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var db *sql.DB

type server struct {
	pb.UnimplementedJuegosServiceServer
}

func (s *server) GetJuegosInfo(ctx context.Context, req *pb.JuegosRequest) (*pb.JuegosResponse, error) {

	var name, win, color string

	query := "select * from chess.tablero where Name like @Name"
	row := db.QueryRowContext(ctx, query, sql.Named("Name", "%"+req.Name+"%"))
	err := row.Scan(&name, &win, &color)

	if err != nil {

		if err == sql.ErrNoRows {
			return &pb.JuegosResponse{

				Name:  "Not found",
				Win:   "Not found",
				Color: "Not found",
			}, nil
		}

		return nil, err
	}
	return &pb.JuegosResponse{
		Name:  name,
		Win:   win,
		Color: color,
	}, nil

}

func (s *server) GetJuegosList(req *pb.Empty, stream pb.JuegosService_GetJuegosListServer) error {

	query := "select * from chess.tablero"
	rows, err := db.Query(query)

	if err != nil {

		log.Panic(err)
		return err

	}

	defer rows.Close()
	for rows.Next() {

		var name, win, color string
		if err := rows.Scan(&name, &win, &color); err != nil {
			log.Panic(err)
			return err
		}

		if err := stream.Send(&pb.JuegosResponse{

			Name:  name,
			Win:   win,
			Color: color,
		}); err != nil {

			log.Panic(err)
			return err
		}
	}
	return nil
}

func (s *server) AddJuegos(stream pb.JuegosService_AddJuegosServer) error {

	var count int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AddJuegosResponse{

				Count: count,
			})
		}

		if err != nil {

			log.Panic(err)
			return err

		}
		query := "insert into chess.tablero (Name, Win, Color) values (@Name, @Win, @Color)"
		_, err = db.Exec(query,
			sql.Named("Name", req.GetName()),
			sql.Named("Win", req.GetWin()),
			sql.Named("Color", req.GetColor()))

		if err != nil {
			log.Panic(err)
			return err
		}

		count++
		log.Printf("Added %s", req.Name)
	}
}

func (s *server) GetJuegosByWin(stream pb.JuegosService_GetJuegosByWinServer) error {

	for {

		req, err := stream.Recv()
		if err == io.EOF {

			log.Printf("End of stream")
			return nil
		}

		if err != nil {

			log.Panic(err)
			return err

		}

		query := "select * from chess.tablero where lower(Win) = lower(@Win)"
		rows, err := db.Query(query, sql.Named("Win", req.Win))
		if err != nil {
			log.Panic(err)
			return err
		}

		defer rows.Close()

		for rows.Next() {

			var name, win, color string

			if err := rows.Scan(&name, &win, &color); err != nil {
				log.Panic(err)
				return err
			}

			if err := stream.Send(&pb.JuegosResponse{

				Name:  name,
				Win:   win,
				Color: color,
			}); err != nil {

				log.Panic(err)
				return err
			}

		}
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, pass, s, port, database)

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterJuegosServiceServer(grpcServer, &server{})

	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
		log.Println("Starting health check server on port :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	log.Println("Starting server on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
