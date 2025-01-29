"use client";
import {useState, useEffect} from "react";
import {useRouter} from "next/navigation";
import {Button, Card, TextInput, Text, Container, Title, Group, Space} from "@mantine/core";
import '@mantine/core/styles.css';
import Cookies from "js-cookie";

export default function AuthPage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [token, setToken] = useState<string | null>(null);
    const router = useRouter();

    useEffect(() => {
        const storedToken = Cookies.get("jwt");
        if (storedToken) {
            setToken(storedToken);
        }
    }, []);

    const handleLogin = async () => {
        const response = await fetch("/api/auth/login", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({username, password})
        });

        if (response.ok) {
            const data = await response.json();
            Cookies.set("jwt", data.token, {expires: 7});
            setToken(data.token);
        }
    };

    const handleLogout = () => {
        Cookies.remove("jwt");
        setToken(null);
    };

    if (token) {
        return (
            <Container size="sm" mt={50}>
                <Card shadow="sm" p="lg" radius="md" withBorder>
                    <Text size="lg">Welcome, {username || "User"}!</Text>
                    <Button color="red" mt={20} onClick={handleLogout}>Logout</Button>
                </Card>
            </Container>
        );
    }

    return (
        <Container size="sm" mt={50}>
            <Card shadow="sm" p="lg" radius="md" withBorder>
                <Title align="center" order={2} mb="lg">Login or Register</Title>
                <TextInput
                    label="Username"
                    placeholder="Enter your username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <TextInput
                    label="Password"
                    placeholder="Enter your password"
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    mt={10}
                />
                <Space h="md"/>
                <Button variant="filled" color="blue" size="xl" radius="xl" onClick={handleLogin}>Login</Button>
            </Card>
        </Container>
    );
}
