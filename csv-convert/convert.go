package csv_convert

//func csvConvert(articles []Articles) {
//	// initializing csv file
//	file, err := os.Create("products.csv")
//	if err != nil {
//		log.Fatalln("failed to create output cSV file", err)
//	}
//
//	defer file.Close()
//	// initializing a file writer
//	writer := csv.NewWriter(file)
//
//	// defining the CSV headers
//	headers := []string{
//		"title",
//		"author",
//		"url",
//	}
//
//	writer.Write(headers)
//
//	// adding each Pokemon product to the CSV output file
//	for _, article := range articles {
//		// converting a Pokemonproduct to an array of strings
//		record := []string{
//			article.title,
//			article.author,
//			article.url,
//		}
//		// writing a new CSV record
//		writer.Write(record)
//	}
//	defer writer.Flush()
//}
