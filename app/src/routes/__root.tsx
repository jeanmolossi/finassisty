import * as React from "react";
import { Link, Outlet, createRootRoute } from "@tanstack/react-router";

export const Route = createRootRoute({
    component: RootComponent,
});

/**
 * Renders the root layout with a navigation bar and an outlet for nested routes.
 *
 * Displays navigation links to the main sections of the application and renders child route components via the outlet.
 */
function RootComponent() {
    return (
        <React.Fragment>
            <nav className="p-4 flex gap-4 bg-gray-100">
                <Link to="/">Home</Link>
                <Link to="/about">About</Link>
                <Link to="/dashboard">Dashboard</Link>
                <Link to="/in-touch">In Touch</Link>
            </nav>

            <Outlet />
        </React.Fragment>
    );
}
