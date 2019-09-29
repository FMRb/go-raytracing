# go-raytracing
Coding a ray tracing 

## Following the introctions from Peter Shirley
[Ray tracing in one weekend](http://www.realtimerendering.com/raytracing/Ray%20Tracing%20in%20a%20Weekend.pdf)

## Resources to understand Ray-tracing
### Video explaining steps in detail
[youtube video](https://www.youtube.com/watch?v=Pc1TbuX2NHk)

### Explanation of dot product
 1. [Vector calculus](https://betterexplained.com/articles/vector-calculus-understanding-the-dot-product/)

### Sphere intersection
[Example intersection ray-sphere](http://kylehalladay.com/blog/tutorial/math/2013/12/24/Ray-Sphere-Intersection.html)
[Analytic solution showing the math](https://www.scratchapixel.com/lessons/3d-basic-rendering/minimal-ray-tracer-rendering-simple-shapes/ray-sphere-intersection)
### Reflection and Refraction
[Vector calculation](https://graphics.stanford.edu/courses/cs148-10-summer/docs/2006--degreve--reflection_refraction.pdf)
    - (20a) formula is plus because of the direction of the normal, however, the program consider the value of "i - cos(theta)n" as sign for the direction of the
        normal is passed by the method calling the refract function.
    - (20b) Appy Pitagoras and use "n" as the direction of the vector.
[Ray tracer refraction with medium](https://computergraphics.stackexchange.com/questions/4573/refraction-in-a-ray-tracer-what-do-with-an-intersection-within-the-medium)