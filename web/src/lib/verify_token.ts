import axios from "axios";

export async function verifyToken(): Promise<App.User | null> {
    const baseApiUrl = import.meta.env.VITE_API_URL;

    try {
        const response = await axios.get(`${baseApiUrl}/me`, {withCredentials: true});

        return response.data;
    } catch (error: any) {
        console.error('error', error.response.data);
        return null;
    }
}
