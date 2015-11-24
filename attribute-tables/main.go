package main

import (
	"fmt"
	"github.com/jonas-p/go-shp"
	"log"
	"reflect"
)

func main() {

	fmt.Println("Sinar Project Attribute Tables Analysis!!")
	// Open shaepfile for reading
	// Needed to somehow cleanup the shapefile!!
	//  /Library/Frameworks/GDAL.framework/Programs/ogr2ogr new-srw-shape.shp srw-shape.shp
	shape, err := shp.Open("/Users/leow/Desktop/TINDAK_MALAYSIA/temp/new-srw-shape.shp")
	if err != nil {
		log.Fatal(err)
	}
	defer shape.Close()

	// fields from attribute tables (DBF)
	fields := shape.Fields()
	for shape.Next() {
		n, p := shape.Shape()

		// print Feature
		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())

		// print Attributes

		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Printf("\t%v : %v\n", f, val)
		}

			fmt.Println(n)
			
			fmt.Println("Boo!!!")
	fmt.Println("ss")
			
	}

}
