import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import App from "./App";
import Market from "./pages/Market";
import Orders from "./pages/Orders";

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/market" element={<Market />} />
        <Route path="/orders" element={<Orders />} />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);
