"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Button, Card, TextInput, Text, Container, Title, Group, Space } from "@mantine/core";
import '@mantine/core/styles.css';
import Link from "next/link";
import Cookies from "js-cookie";
import ErrorNotification from "../components/ErrorNotification"; // Importing the ErrorNotification component

export default function RegisterPage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const router = useRouter();

    useEffect(() => {
        const refreshToken = Cookies.get("refresh");
        if (refreshToken) {
            router.push("/");
        }
    }, [router]);

    const handleRegister = async () => {
        const response = await fetch("http://localhost:8081/user/signUp", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: 'include',
            body: JSON.stringify({ username, password }),
        });

        if (response.ok) {
            router.push("/sign-in");
        } else {
            const responseMessage = await response.json();
            setErrorMessage(responseMessage.error || "Registration failed");
            setTimeout(() => setErrorMessage(null), 7000);
        }
    };

    return (
        <>
            <ErrorNotification errorMessage={errorMessage} onClose={() => setErrorMessage(null)} />

            <Container size="sm" mt={50}>
                <Card shadow="sm" p="lg" radius="md" withBorder>
                    <Title ta="center" order={2} mb="lg">Sign Up</Title>
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
                    <Space h="md" />
                    <Button variant="filled" color="blue" size="xl" radius="xl" onClick={handleRegister}>
                        Sign Up
                    </Button>
                    <Space h="md" />
                    <Group justify="center" mt="md">
                        <Text size="sm">Already have an account?</Text>
                        <Link href="/sign-in" passHref>
                            <Text size="sm" td="underline"
                                  style={{ cursor: "pointer", fontWeight: "bold", color: "black" }}>
                                Sign In
                            </Text>
                        </Link>
                    </Group>
                </Card>
            </Container>
        </>
    );
}
