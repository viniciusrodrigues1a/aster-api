package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"product-service/cmd/product-service/application/use-case"
	"product-service/cmd/product-service/domain/dto"
	"strconv"
)

type FieldMustBeANumberError struct {
	fieldName string
}

func (f *FieldMustBeANumberError) Error() string {
	return fmt.Sprintf("Field %s must be a number", f.fieldName)
}

type CreateProductController struct {
	useCase *usecase.CreateProductUseCase
}

func NewCreateProductController(useCase *usecase.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		useCase: useCase,
	}
}

func (c *CreateProductController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		fieldErr := FieldMustBeANumberError{fieldName: "quantity"}
		http.Error(w, fieldErr.Error(), http.StatusBadRequest)
	}

	purchasePrice, err := strconv.Atoi(r.FormValue("purchase_price"))
	if err != nil {
		fieldErr := FieldMustBeANumberError{fieldName: "purchase_price"}
		http.Error(w, fieldErr.Error(), http.StatusBadRequest)
	}

	salePrice, err := strconv.Atoi(r.FormValue("sale_price"))
	if err != nil {
		fieldErr := FieldMustBeANumberError{fieldName: "sale_price"}
		http.Error(w, fieldErr.Error(), http.StatusBadRequest)
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image")
	var image *dto.ProductImage
	var buf *bytes.Buffer
	if err == nil {
		defer file.Close()
		buf = bytes.NewBuffer(nil)
		io.Copy(buf, file)

		image = &dto.ProductImage{
			Base64:   base64.StdEncoding.EncodeToString(buf.Bytes()),
			MimeType: handler.Header.Get("Content-Type"),
		}
	}

	useCaseErr := c.useCase.Execute(&usecase.CreateProductUseCaseRequest{
		Title:         r.FormValue("title"),
		Description:   r.FormValue("description"),
		Quantity:      int32(quantity),
		PurchasePrice: int64(purchasePrice),
		SalePrice:     int64(salePrice),
		AccountID:     r.Context().Value("account_id").(string),
		Image:         image,
	})
	if useCaseErr != nil {
		status := http.StatusInternalServerError

		if useCaseErr == usecase.ErrTitleIsEmpty ||
			useCaseErr == usecase.ErrPurchasePriceZero ||
			useCaseErr == usecase.ErrSalePriceZero {
			status = http.StatusBadRequest
		}

		http.Error(w, useCaseErr.Error(), status)
	}
}
