import React from "react";
import TruckList from "../components/TruckList";
import OrderSummary from "../components/Order";

const CustomerPage = () => {
  return (
    <div className="w-full flex flex-col items-center justify-start h-screen">
      <h1 className="text-3xl font-bold mb-4">Customer Page</h1>
      <div className="w-full overflow-y-auto flex p-4 gap-4">
        <TruckList role="customer" />
        <OrderSummary role="customer" />
      </div>
    </div>
  );
};

export default CustomerPage;
