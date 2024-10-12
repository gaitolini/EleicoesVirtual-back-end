package services

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// FirestoreClientInterface define uma interface para o cliente Firestore
// Isso permite que a interface seja mockada durante os testes
type FirestoreClientInterface interface {
	Collection(path string) FirestoreCollectionInterface
}

// FirestoreCollectionInterface define uma interface para a coleção Firestore
// Isso permite que a coleção seja mockada durante os testes
type FirestoreCollectionInterface interface {
	Doc(docID string) FirestoreDocumentRefInterface
	Documents(ctx context.Context) FirestoreDocumentIteratorInterface
	Add(ctx context.Context, data interface{}) (FirestoreDocumentRefInterface, *firestore.WriteResult, error)
}

// FirestoreDocumentRefInterface define uma interface para o documento Firestore
// Isso permite que o documento seja mockado durante os testes
type FirestoreDocumentRefInterface interface {
	Set(ctx context.Context, data interface{}) (*firestore.WriteResult, error)
	Get(ctx context.Context) (*firestore.DocumentSnapshot, error)
	Delete(ctx context.Context) (*firestore.WriteResult, error)
}

// FirestoreDocumentIteratorInterface define uma interface para o iterador de documentos Firestore
// Isso permite que o iterador seja mockado durante os testes
type FirestoreDocumentIteratorInterface interface {
	Next() (*firestore.DocumentSnapshot, error)
}

// FirestoreClientWrapper é a implementação padrão da interface FirestoreClientInterface
type FirestoreClientWrapper struct {
	Client *firestore.Client
}

// Collection retorna uma referência à coleção especificada
func (fc *FirestoreClientWrapper) Collection(path string) FirestoreCollectionInterface {
	return &FirestoreCollectionWrapper{Collection: fc.Client.Collection(path)}
}

// FirestoreCollectionWrapper é a implementação padrão da interface FirestoreCollectionInterface
type FirestoreCollectionWrapper struct {
	Collection *firestore.CollectionRef
}

// Doc retorna uma referência ao documento especificado
func (fc *FirestoreCollectionWrapper) Doc(docID string) FirestoreDocumentRefInterface {
	return &FirestoreDocumentRefWrapper{DocumentRef: fc.Collection.Doc(docID)}
}

// Documents retorna um iterador para os documentos na coleção
func (fc *FirestoreCollectionWrapper) Documents(ctx context.Context) FirestoreDocumentIteratorInterface {
	return &FirestoreDocumentIteratorWrapper{DocumentIterator: fc.Collection.Documents(ctx)}
}

// Add adiciona um novo documento à coleção
func (fc *FirestoreCollectionWrapper) Add(ctx context.Context, data interface{}) (FirestoreDocumentRefInterface, *firestore.WriteResult, error) {
	docRef, writeResult, err := fc.Collection.Add(ctx, data)
	if err != nil {
		log.Printf("Erro ao adicionar documento: %v", err)
		return nil, nil, err
	}
	return &FirestoreDocumentRefWrapper{DocumentRef: docRef}, writeResult, nil
}

// FirestoreDocumentRefWrapper é a implementação padrão da interface FirestoreDocumentRefInterface
type FirestoreDocumentRefWrapper struct {
	DocumentRef *firestore.DocumentRef
}

// Set define os dados do documento especificado
func (fd *FirestoreDocumentRefWrapper) Set(ctx context.Context, data interface{}) (*firestore.WriteResult, error) {
	return fd.DocumentRef.Set(ctx, data)
}

// Get retorna os dados do documento especificado
func (fd *FirestoreDocumentRefWrapper) Get(ctx context.Context) (*firestore.DocumentSnapshot, error) {
	return fd.DocumentRef.Get(ctx)
}

// Delete remove o documento especificado
func (fd *FirestoreDocumentRefWrapper) Delete(ctx context.Context) (*firestore.WriteResult, error) {
	return fd.DocumentRef.Delete(ctx)
}

// FirestoreDocumentIteratorWrapper é a implementação padrão da interface FirestoreDocumentIteratorInterface
type FirestoreDocumentIteratorWrapper struct {
	DocumentIterator *firestore.DocumentIterator
}

// Next retorna o próximo documento do iterador
func (fdi *FirestoreDocumentIteratorWrapper) Next() (*firestore.DocumentSnapshot, error) {
	return fdi.DocumentIterator.Next()
}
