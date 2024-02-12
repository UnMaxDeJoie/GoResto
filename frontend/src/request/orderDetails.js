import { API_URL } from "./api";

export const createOrder = async (orderDetailsData) => {
    try {
        const response = await fetch(`${API_URL}/orderorder_detail`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(orderDetailsData),
        });
        if (!response.ok) {
            throw new Error("Error creating orderDetails");
        }
        return await response.json();
    } catch (error) {
        console.error("There was an error creating the orderDetails:", error);
        throw error;
    }
};
export const getOrderDetailsByOrderID = async (orderId) => {
    const response = await fetch(`${API_URL}/order_detail/order/${orderId}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting order Detail by order Id");
    }
    return response.json();
};
export const getOrderDetailsByTruckID = async (truckId) => {
    const response = await fetch(`${API_URL}/order_detail/truck/${truckId}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting order Detail by truck Id");
    }
    return response.json();
};
export const getOrderDetailsByConsumableID = async (consumableId) => {
    const response = await fetch(`${API_URL}/order_detail/truck/${consumableId}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting order Detail by consumable Id");
    }
    return response.json();
}
export const updateOrderDetail = async (id) => {
    const response = await fetch(`${API_URL}/order_detail/order/${id}`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting order details");
    }
    return response.json();
}
export const deleteOrderDetails = async (orderId, consumableId) => {
    const response = await fetch(`${API_URL}/order_detail/order/${orderId}/consumable/${consumableId}`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting order details");
    }
    return response.json();
}
