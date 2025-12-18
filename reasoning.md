# Design Decisions & Thought Process

This document outlines the reasoning behind the architectural choices I made while building this User Management Service. My main goal wasn't just to make it "work," but to build something that is maintainable, scalable, and easy to test.

Here is a breakdown of why I did what I did.

## 1. Three-Layered Architecture
I decided to split the application into three distinct layers to keep concerns seperate:
* **Repository Layer:** Strictly handles the database interactions (SQLC generated code). It doesn't know about HTTP or business logic.
* **Service Layer:** This is where the actual logic lives (like the Age Calculation). It sits in the middle.
* **Handler Layer:** Manages the HTTP request/response cycle and input parsing.

This separation makes the code much cleaner. If I need to change how the age is calculated, I only touch the Service layer. If I need to change the URL path, I only touch the Handler.

## 2. Dependency Injection (Wiring it all up)
Instead of using global variables (which are a nightmare to test), I used **Dependency Injection** to wire the layers together in `main.go`.
* The `Service` struct requires a `Repository` interface.
* The `Handler` struct requires a `Service` interface.

This means in `main.go`, I explicitely create the connection, create the repository, pass it to the service, and pass that to the handler. It makes the dependency graph very clear.

## 3. Injecting Logger and Validator
I applied the same injection pattern to the **Zap Logger** and the **Validator**.
Instead of initializing them inside the Handler functions, I injected them into the `Handler` struct when the application starts.
* **Why?** It makes the Handler "dumb" in a good way. It doesn't need to know *how* to create a logger, it just knows it has one available to use. This also means I could theoretically pass in a mock logger or validator during testing without breaking anything.

## 4. Database Access (SQLC)
For the database interaction, I chose **SQLC** to generate Go code from raw SQL queries.
* **Reasoning:** I wanted full control over the SQL queries without the risk of typos. SQLC gives me the best of both worlds: I write the SQL, and it generates type-safe Go structs. If I change a column name in the DB but forget to update the Go code, the build actully fails immediately. This type-safety is a huge confidence booster.

## 5. Docker & Multi-Stage Builds
I didn't want to ship a massive container with the full Go compiler inside.
I used a **Multi-Stage Docker build**:
1.  **Stage 1:** Uses the heavy Go image to compile the binary.
2.  **Stage 2:** Copies *only* the binary into a tiny Alpine Linux image.
This reduced the final image size drastically (around 15-20MB), which makes deployments faster and more secure since the source code isn't lying around in the container.

## Conclusion
Overall, I tried to balance "getting it done" with "doing it right." The dependency injection setup took a bit more boilerplate code in `main.go` compared to just using globals, but the resulting stability and clear structure were definately worth it.
