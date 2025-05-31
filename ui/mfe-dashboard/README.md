# MFE Dashboard

Independent micro-frontend that displays progress charts and student statistics:
- Built in Angular 19.
- Publishes a bundle that the shell (`web-app`) consumes via Module Federation.

## How to start in dev mode

```bash
  cd ui/mfe-dashboard
```
```bash
  npm install
```
```bash
  ng serve --port 4300
```

## Deploy
- `npm run build --prod` → generates `dist/`
- Copy dist to the Nginx container configured in Dockerfile.

## Best practices
- Keep CSS scoped in the component.
- Version the publicPathto break the cache when releasing a new version.
