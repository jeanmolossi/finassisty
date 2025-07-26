import React, { Suspense } from "react";
import { Link, Outlet } from "react-router";

export default function App() {
    return (
        <>
            <nav className="p-4 flex gap-4 bg-gray-100">
                <Link to="/">Home</Link>
                <Link to="/about">About</Link>
                <Link to="/dashboard">Dashboard</Link>
                <Link to="/in-touch">In Touch</Link>
            </nav>
            <Suspense fallback="Loading...">
                <Outlet />
            </Suspense>
        </>
    );
}
