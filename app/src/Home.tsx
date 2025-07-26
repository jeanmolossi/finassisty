import { Switch } from "@base-ui-components/react/switch";
import { useLoaderData } from "react-router";

export async function loader() {
    await fetch("http://localhost:8080/api/hello", {
        method: "GET",
    });
    return { message: "Hello, world!" };
}

export default function Home() {
    const loaderData = useLoaderData();
    console.log(loaderData);

    return (
        <main className="p-4">
            <h1 className="text-2xl font-bold mb-4">Finassisty Home 2</h1>
            <Switch.Root defaultChecked>
                <Switch.Thumb />
            </Switch.Root>
        </main>
    );
}
