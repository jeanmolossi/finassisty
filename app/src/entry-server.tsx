import { createRouter } from '@tanstack/react-router';
import { RouterServer } from '@tanstack/react-router/ssr/server';
import { routeTree } from './routeTree.gen';

export default function entryServer() {
  const router = createRouter({
    routeTree,
    defaultPreload: 'intent',
  }) as any; // eslint-disable-line @typescript-eslint/no-explicit-any
  return <RouterServer router={router} />;
}
