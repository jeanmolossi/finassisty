import { aboutRoute } from "../../router/routes";
import { Switch } from "@base-ui-components/react/switch";

export default function Home() {
    const loaderData = aboutRoute.useLoaderData();

    return (
        <main className="p-4">
            <h1 className="text-2xl font-bold mb-4">Finassisty About</h1>
            <Switch.Root defaultChecked>
                <Switch.Thumb />
            </Switch.Root>

            <pre>
                <code>{JSON.stringify(loaderData)}</code>
            </pre>
        </main>
    );
}
