package main

import "fmt"

type product struct {
	title string
	price string
}

type detail struct {
	image   string
	content string
}

type productWithDetail struct {
	test string
	product 
	detail  
}

func NewProduct(title, price string) product {
	return product{
		title: title,
		price: price,
	}
}

func NewDetail(image, content string) detail {
	return detail{
		image:   image,
		content: content,
	}
}

func ProductWithDetail(p product, d detail) productWithDetail {
	return productWithDetail{
		p ,d
	}
}

func (p productWithDetail) display() {
	fmt.Print(p.product)
	// fmt.Printf("Name: %v\nPrice: %v\nImage: %v\nContent: %v", p.title, p.price, p.image, p.content)
}

func main() {
	var product = NewProduct("Car", "12")
	var detail = NewDetail("Image", "Fast")
	var newProductWithDetail = ProductWithDetail(product, detail)
	newProductWithDetail.display()
}
