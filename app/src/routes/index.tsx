import { Switch } from "@base-ui-components/react/switch";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({ component: Home });

/**
 * Renders the home page with a heading and a toggle switch UI.
 *
 * Displays the "Finassisty Home" title and a switch initialized in the checked state.
 */
function Home() {
    return (
        <main className="p-4">
            <h1 className="text-2xl font-bold mb-4">Finassisty Home</h1>
            <Switch.Root defaultChecked>
                <Switch.Thumb />
            </Switch.Root>
        </main>
    );
}
