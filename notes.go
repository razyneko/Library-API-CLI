package main

//  go mod init example/(folder name)
// usually a path where to download this package
// go.mod keeps track of dependencies of the project

// we add `json :"--"` for serialising struct to conver to json because apis work by sending and recieving json --> this just says that in json convert it to that
// we use capital keys in structs because that makes it an exported ,field a public field which can be viewed by outer modules
// if the first letter is small in struct keys .. we will empty json in return every time because it isnt read
// *gin.Context all of info about request allows to return a response
// *gin.Context stores all info related to specific request (query params data payload headers)
// to bind json which was part of data payload of request to the book struct
// with pointer .. direct modification of field values can be done
// .BindJSON() method is what will handle sending the error response
// reference is given to modify fields of data structure from a diff func
