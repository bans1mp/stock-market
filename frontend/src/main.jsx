import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import App from "./App";
import Market from "./pages/Market";
import Orders from "./pages/Orders";
import Login from "./pages/Login";
import OrderPage from "./pages/OrderPage";
import Register from "./pages/Register";
import StockDetail from "./pages/StockDetail";
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';
import Stocks from "./pages/Stocks";


ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/stocks" element={<Stocks />} />
        <Route path="/market" element={<Market />} />
        <Route path="/orders" element={<Orders />} />
        <Route path="/buy-order/:symbol" element={<OrderPage orderType="buy" />} />
        <Route path="/sell-order/:symbol" element={<OrderPage orderType="sell" />} />
        <Route path="/stocks/:stockSymbol" element={<StockDetail />} />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);
