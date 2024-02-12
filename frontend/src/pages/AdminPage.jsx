import React from "react";

import OrderSummary from "../components/Order";
import TruckList from "../components/TruckList";

const AdminPage = () => {
  return (
    <div className="w-full flex flex-col items-center justify-center h-screen overflow-y-auto">
      <h1 className="text-3xl font-bold mb-4">Admin Page</h1>
      <div className="w-full overflow-y-auto flex p-4 gap-4">
        <TruckList role="admin" />
        <OrderSummary role="admin" />
      </div>
      <div className="w-full flex flex-col items-center justify-center gap-4"></div>
    </div>
  );
};

export default AdminPage;
