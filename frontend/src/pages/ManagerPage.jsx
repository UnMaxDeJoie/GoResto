import React from "react";
import OrderSummary from "../components/Order";
import TruckList from "../components/TruckList";
import CreateTruck from "../components/CreateTruck";
import EditTruckForm from "../components/EditTruckForm";

const ManagerPage = () => {
  const [selectedTruckData, setSelectedTruckData] = React.useState(null);

  const handleSaveTruck = async (truckData) => {
    const editedTruckData = {
      name: truckData.name,
      slotBuffer: parseInt(truckData.slotBuffer, 10),
      opening: truckData.opening,
      closing: truckData.closing,
      userID: parseInt(truckData.userID, 10),
    };
    console.log(editedTruckData);

    // TODO: Appeler la fonction d'édition de truck
  };

  return (
    <div className="w-full flex flex-col items-center justify-center h-screen overflow-y-auto">
      <h1 className="text-3xl font-bold mb-4">Manager Page</h1>
      <div className="w-full overflow-y-auto flex p-4 gap-4">
        <TruckList role="manager" />
        <OrderSummary role="manager" />
      </div>
      <div className="w-full flex flex-col items-center justify-center gap-4">
        {/* <CreateTruck /> */}
        {/* 
            // edit truck form
        {selectedTruckData && (
          <EditTruckForm
            selectedTruckData={selectedTruckData}
            onSave={handleSaveTruck}
          />
        )} */}
      </div>
    </div>
  );
};

export default ManagerPage;
