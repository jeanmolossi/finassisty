import { Switch } from "@base-ui-components/react/switch";
import { createFileRoute } from "@tanstack/react-router";
import loader from "./loader";

export const Route = createFileRoute("/about/")({
    component: AboutPage,
    wrapInSuspense: true,
    loader: async () => loader(),
});

function AboutPage() {
    const loaderData = Route.useLoaderData();

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
