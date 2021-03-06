// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"cloud.google.com/go/datastore"
	"github.com/juntaki/techbook-qrcode/src/application"
	"github.com/juntaki/techbook-qrcode/src/infra"
)

// Injectors from wire.go:

func InitializeQRCodeServiceServer(db *datastore.Client) *application.QRCodeServiceServer {
	qrCodeRepository := infra.NewQRCodeRepositoryDatastoreImpl(db)
	techBookRepository := infra.NewTechBookRepositoryDatastoreImpl(db)
	qrCodeServiceServer := application.NewQRCodeServiceServer(qrCodeRepository, techBookRepository)
	return qrCodeServiceServer
}

func InitializeTechBookServer(db *datastore.Client) *application.TechBookServer {
	techBookRepository := infra.NewTechBookRepositoryDatastoreImpl(db)
	qrCodeRepository := infra.NewQRCodeRepositoryDatastoreImpl(db)
	techBookServer := application.NewTechBookServer(techBookRepository, qrCodeRepository)
	return techBookServer
}
