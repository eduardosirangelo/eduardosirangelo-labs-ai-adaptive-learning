# Web App (Shell)

Main host for micro-frontends, based on Angular 19 + Module Federation:
- Dynamically loads remotes (e.g., dashboard, reports).
- Manages routing, layout, and global authentication.

## How to start

```bash
  cd ui/web-app
```
```bash
  npm install
```
```bash
  ng serve --port 4200
```

## Configuration points
- `webpack.config.js` → `ModuleFederationPlugin`
- `environment.ts` → URL of remotes and API gateway

## To add a new MFE
- Register the name and URL of the remote in `webpack.config.js`.
- Create an Angular route to load the remote component.