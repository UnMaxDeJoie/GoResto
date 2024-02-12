import { API_URL } from "./api";

export const createOrder = async (orderData) => {
    try {
        const response = await fetch(`${API_URL}/order`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(orderData),
        });
        if (!response.ok) {
            throw new Error("Error creating order");
        }
        return await response.json();
    } catch (error) {
        console.error("There was an error creating the order:", error);
        throw error;
    }
};
export const getOrderById = async (id) => {
    const response = await fetch(`${API_URL}/order/${id}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting order");
    }
    return response.json();
}
export const updateOrder = async (id) => {
    const response = await fetch(`${API_URL}/order/${id}`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting order");
    }
    return response.json();
};
export const deleteOrder = async (id) => {
    const response = await fetch(`${API_URL}/order/${id}`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting order");
    }
    return response.json();
};