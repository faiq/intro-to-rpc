//here we are creating our own type, for readability, "image" which is an array of bytes
typedef binary image

service MakeTags {
    // here we are defining a service that has a method generate which returns a list of strings and    // takes in an image
    list<string> generate(1: image img)
}
