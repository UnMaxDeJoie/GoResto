import { useEffect, useState } from "react";

import UserList from "./UserList";

const testOrders = [
  { id: 1, quantity: 10, comment: "A delicious burger", status: "pending" },
  { id: 2, quantity: 5, comment: "A delicious fries", status: "done" },
  { id: 3, quantity: 7, comment: "A delicious salad", status: "declined" },
  { id: 4, quantity: 10, comment: "A delicious pizza", status: "pending" },
  { id: 5, quantity: 5, comment: "A delicious pasta", status: "pending" },
  { id: 6, quantity: 7, comment: "A delicious salad", status: "pending" },
];

const OrderSummary = (props) => {
  const [orders, setOrders] = useState(testOrders);

  const userType = localStorage.getItem("userType");

  useEffect(() => {
    // Fetch orders from the server
  }, []);

  const handleAccept = (orderId) => {
    if (userType === "admin" || userType === "conservator") {
      // Update the order status to "done"
    }
  };

  const handleReject = (orderId) => {
    if (userType === "admin" || userType === "conservator") {
      // Update the order status to "declined"
    }
  };

  return (
    <div className="container mx-auto flex flex-col items-center">
      <div className="container mx-auto">
        <h1 className="text-2xl font-bold mb-4">Order Summary</h1>
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Quantity
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Comment
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Status
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Actions
              </th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {orders.map((order) => (
              <tr key={order.id}>
                <td className="px-6 py-4 whitespace-nowrap">
                  {order.quantity}
                </td>
                <td className="px-6 py-4 whitespace-nowrap">{order.comment}</td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <span
                    className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                      order.status === "pending"
                        ? "bg-yellow-100 text-yellow-800"
                        : order.status === "done"
                        ? "bg-green-100 text-green-800"
                        : "bg-red-100 text-red-800"
                    }`}
                  >
                    {order.status}
                  </span>
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  {order.status === "pending" &&
                    (props.role === "admin" || props.role === "manager") && (
                      <>
                        <button
                          className="text-green-600 hover:text-green-900 mr-2"
                          onClick={() => handleAccept(order.id)}
                        >
                          Accept
                        </button>
                        <button
                          className="text-red-600 hover:text-red-900"
                          onClick={() => handleReject(order.id)}
                        >
                          Reject
                        </button>
                      </>
                    )}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      {props.role === "admin" && <UserList />}
    </div>
  );
};

export default OrderSummary;
