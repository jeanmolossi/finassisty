export default async function loader() {
    const response = await fetch("http://localhost:8080/api/v1/hello", {
        method: "GET",
    });

    return response.json();
}
