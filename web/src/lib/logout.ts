import { goto } from "$app/navigation";
import axios from "axios";

export async function logout() {
    const baseApiUrl = import.meta.env.VITE_API_URL;

    // clear local storage
    localStorage.removeItem('jwt');

    // call api to logout
    try {
        const response = await axios.post(
            `${baseApiUrl}/logout`,
            {},
            {
                withCredentials: true,
            }
        );

        if (response.status !== 200) {
            console.error('Failed to logout');
        }
    } catch (error) {
        console.error(error);
    }

    // redirect to main page
    goto('/');
}
