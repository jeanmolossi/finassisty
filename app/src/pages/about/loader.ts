export default async function loader() {
    await fetch("http://localhost:8080/api/hello", {
        method: "GET",
    });

    return { message: "Hello, world!" };
}
