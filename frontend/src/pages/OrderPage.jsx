import React, { useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import Navbar from "../components/Navbar.jsx";
import Footer from "../components/Footer.jsx";
import homeImage from "../assets/home.jpg";

const OrderPage = ({ orderType }) => {
  const { symbol } = useParams();
  const [quantity, setQuantity] = useState("");
  const [price, setPrice] = useState("");

  const placeOrder = async () => {
    try {
      const token = localStorage.getItem("token");
      if (!token) {
        alert("Unauthorized: No token found");
        return;
      }

      const orderData = {
        user_id: 1,
        symbol,
        price: parseFloat(price),
        quantity: parseInt(quantity),
        order_type: orderType,
      };

      const response = await axios.post("http://localhost:8080/" + `${orderType}`, orderData, {
        headers: { Authorization: token },
      });
      console.log(response)
      alert(response.data.message);
    } catch (error) {
      console.error("Error placing order:", error);
      alert("Failed to place order");
    }
  };

  return (
    <div
      className="d-flex flex-column min-vh-100 text-white"
      style={{
        backgroundImage: `url(${homeImage})`,
        backgroundSize: "cover",
        backgroundPosition: "center",
        backgroundRepeat: "no-repeat",
        backgroundBlendMode: "overlay",
        backgroundColor: "rgba(0, 0, 0, 0.85)",
      }}
    >
      <Navbar />
      <div className="container flex-grow-1 d-flex flex-column align-items-center justify-content-center mt-5 px-5">
        <h1
          className="fw-bold display-2 text-success border-start border-4 border-success ps-3 mb-4"
          style={{ fontFamily: "Poppins, sans-serif" }}
        >
          {orderType === "buy" ? "Buy" : "Sell"} <span className="text-danger">{symbol}</span>
        </h1>
        <div className="card bg-dark text-white p-4 rounded shadow w-50 text-center">
          <h2 className="mb-4">Place {orderType} Order</h2>
          <p className="fs-5">Stock: <strong>{symbol}</strong></p>
          <div className="input-group mb-3">
            <span className="input-group-text bg-success text-white border-0" style={{ fontWeight: "bold" }}>
              Qty
            </span>
            <input
              type="number"
              placeholder="Enter quantity"
              className="form-control text-center bg-dark text-white border-success shadow-sm"
              style={{ transition: "0.3s", outline: "none" }}
              onFocus={(e) => (e.target.style.borderColor = "#28a745")}
              onBlur={(e) => (e.target.style.borderColor = "#6c757d")}
              value={quantity}
              onChange={(e) => setQuantity(e.target.value)}
            />
          </div>
          <div className="input-group mb-3">
            <span className="input-group-text bg-warning text-dark border-0" style={{ fontWeight: "bold" }}>
              Price
            </span>
            <input
              type="number"
              placeholder="Enter price"
              className="form-control text-center bg-dark text-white border-warning shadow-sm"
              style={{ transition: "0.3s", outline: "none" }}
              onFocus={(e) => (e.target.style.borderColor = "#ffc107")}
              onBlur={(e) => (e.target.style.borderColor = "#6c757d")}
              value={price}
              onChange={(e) => setPrice(e.target.value)}
            />
          </div>
          <button className={`btn btn-${orderType === "buy" ? "success" : "danger"} w-100`} onClick={placeOrder}>
            Confirm {orderType} Order
          </button>
        </div>
      </div>
      <Footer />
    </div>
  );
};

export default OrderPage;
