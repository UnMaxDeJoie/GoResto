import { API_URL } from "./api";

export const Signin = async (email, password) => {
    const response = await fetch(`${API_URL}/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
    });
    if (!response.ok) {
        throw new Error("Error logging in");
    }
    return response.json();
};

export const Signup = async (userName, email, password, userRole) => {
    const response = await fetch(`${API_URL}/register`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ username: userName, email, password, role: userRole }),
    });
    if (!response.ok) {
        throw new Error("Error registering");
    }
    return response.json();
};

export const GetUserById = async (id) => {
    const response = await fetch(`${API_URL}/user/${id}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting user");
    }
    return response.json();
}

export const deleteAccount = async (userID) => {
    const response = await fetch(`${API_URL}/users/${userID}`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error deleting user");
    }
    console.log(response.json());
    return response.json();
};
