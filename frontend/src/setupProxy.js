const proxy = require("http-proxy-middleware");

module.exports = (app) => {
  app.use(
    proxy.createProxyMiddleware("/api", { target: "http://localhost:8000/" })
  );
};
