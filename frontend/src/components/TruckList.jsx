import { useState, useEffect } from "react";
import { NavLink } from "react-router-dom";

import { GetAvailableTrucks, getTruckById } from "./../request/truck";

const testTruckData = [
  {
    id: 1,
    name: "Restaurant 1",
    opening: "10:00",
    closing: "22:00",
    slotBuffer: 10,
    menu: [
      {
        label: "Burger",
        price: 10,
        description: "A delicious burger",
      },
      {
        label: "Fries",
        price: 5,
        description: "A delicious fries",
      },
      {
        label: "Salad",
        price: 7,
        description: "A delicious salad",
      },
    ],
  },
  {
    id: 2,
    name: "Restaurant 2",
    opening: "10:00",
    closing: "22:00",
    slotBuffer: 10,
    menu: [
      {
        label: "Pizza",
        price: 10,
        description: "A delicious pizza",
      },
      {
        label: "Pasta",
        price: 5,
        description: "A delicious pasta",
      },
      {
        label: "Salad",
        price: 7,
        description: "A delicious salad",
      },
    ],
  },

  {
    id: 3,
    name: "Restaurant 3",
    opening: "10:00",
    closing: "22:00",
    slotBuffer: 10,
    menu: [
      {
        label: "Sushi",
        price: 10,
        description: "A delicious sushi",
      },
      {
        label: "Ramen",
        price: 5,
        description: "A delicious ramen",
      },
    ],
  },

  {
    id: 4,
    name: "Restaurant 4",
    opening: "10:00",
    closing: "22:00",
    slotBuffer: 10,
    menu: [
      {
        label: "Tacos",
        price: 10,
        description: "A delicious tacos",
      },
      {
        label: "Burrito",
        price: 5,
        description: "A delicious burrito",
      },
    ],
  },
];

const TruckList = (props) => {
  const [currentTruck, setCurrentTruck] = useState(null);
  const [truckData, setTruckData] = useState([]);

  const handleTruckClick = (truck) => {
    setCurrentTruck(truck);
  };

  //   useEffect(() => {
  //     const fetchTrucks = async () => {
  //       try {
  //         // récupère les trucks disponibles
  //         const trucks = await GetAvailableTrucks();
  //         setTruckData(trucks);
  //       } catch (error) {
  //         console.error("Failed to fetch trucks", error);
  //       }
  //     };

  //     fetchTrucks();
  //     if (role === "manager") {
  //       // garde les trucks ayant un userID correspondant à celui de l'utilisateur
  //       const filteredTrucks = truckData.filter(
  //         (truck) => truck.userID === localStorage.getItem("id")
  //       );
  //       setTruckData(filteredTrucks);
  //     }
  //   }, []);

  const handleDelete = async (id) => {
    try {
      await getTruckById(id);
      console.log("Truck deleted successfully, please refresh the page");
    } catch (error) {
      console.error("There was an error deleting the truck:", error);
    }
  };

  return (
    <div className="flex flex-col h-full w-full p-4 gap-4">
      <span>
        <h1 className="text-2xl font-bold mb-4">Truck List</h1>{" "}
        <span>(Click on a truck to see its menu)</span>
      </span>
      <div className="w-full h-1/5 overflow-y-auto border border-gray-200">
        {testTruckData.map((testTruckData) => (
          <div
            key={testTruckData.id}
            className="bg-white p-4 shadow hover:bg-gray-100 border-b border-gray-200 cursor-pointer"
            onClick={() => handleTruckClick(testTruckData)}
          >
            <h2 className="text-lg font-bold">{testTruckData.name}</h2>
            <p>{testTruckData.hours}</p>
          </div>
        ))}
      </div>
      {currentTruck && (
        <div className="w-full bg-white p-4 flex flex-col justify-between items-start gap-4">
          <div>
            <h2 className="text-lg font-bold">Selected Truck:</h2>
            <p>{currentTruck.name}</p>
            <p>
              opening: {currentTruck.opening} - closing: {currentTruck.closing}
            </p>

            <p> Slot Buffer: {currentTruck.slotBuffer}</p>
          </div>
          <h2 className="text-lg font-bold">Menu:</h2>
          <div className="w-full flex flex-col gap-2 overflow-y-auto max-h-64">
            {currentTruck &&
              currentTruck.menu.map((item) => (
                <div
                  key={item.label}
                  className="bg-gray-100 p-4 shadow rounded"
                >
                  <h3 className="text-md font-bold">{item.label}</h3>
                  <p>Price: ${item.price}</p>
                  <p>Description: {item.description}</p>
                  <div className="flex flex-row items-center gap-2">
                    {props.role === "customer" && (
                      <button
                        disabled={props.role !== "customer"}
                        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                      >
                        Order Now
                      </button>
                    )}
                    {props.role === "manager" ||
                      (props.role === "admin" && (
                        <button
                          className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                          disabled={props.role === "customer"}
                          onClick={() => handleDelete(currentTruck.id)}
                        >
                          Delete
                        </button>
                      ))}
                    {props.role === "manager" && (
                      <button
                        className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                        disabled={props.role === "customer"}
                      >
                        Edit
                      </button>
                    )}
                  </div>
                </div>
              ))}
          </div>
        </div>
      )}
    </div>
  );
};

export default TruckList;
