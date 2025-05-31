package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Cliente *mongo.Client
var PostosCollection *mongo.Collection

func ConectarMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Múltiplos endereços para conexão ao cluster sharded
	servidores := []string{
		"mongodb://servidor1:27017",
		"mongodb://servidor2:27017",
		"mongodb://servidor3:27017",
	}

	// Opções de conexão com suporte para replicação e sharding
	clienteOpts := options.Client().
		SetHosts(servidores).
		SetReplicaSet("meuReplicaSet"). // Defina o nome do seu replica set
		SetRetryWrites(true).           // Habilita a tentativa de operações de escrita
		SetRetryReads(true).            // Habilita a tentativa de operações de leitura
		SetServerSelectionTimeout(5 * time.Second)

	client, erro := mongo.Connect(ctx, clienteOpts)
	if erro != nil {
		return erro
	}

	Cliente = client

	// Conecta ao banco de dados e coleção
	PostosCollection = client.Database("reservasRedes").Collection("postos")
	return nil
}
