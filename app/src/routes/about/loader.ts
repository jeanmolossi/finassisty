/**
 * Fetches and returns JSON data from the local API endpoint `/api/v1/hello`.
 *
 * @returns The parsed JSON response from the API.
 */
export default async function loader() {
    const response = await fetch("http://localhost:8080/api/v1/hello", {
        method: "GET",
    });

    return response.json();
}
