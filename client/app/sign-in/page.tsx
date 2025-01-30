"use client";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Button, Card, TextInput, Text, Container, Title, Group, Space, Notification } from "@mantine/core";
import '@mantine/core/styles.css';
import Link from "next/link";
import { IconX } from '@tabler/icons-react';
import Cookies from "js-cookie";

export default function LoginPage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const router = useRouter();

    useEffect(() => {
        // Check if JWT token exists in cookies, if yes, redirect to the main page
        const storedToken = Cookies.get("jwt");
        if (storedToken) {
            router.push("/"); // If token is found, redirect to the main page
        }
    }, [router]);

    const handleSignIn = async () => {
        const response = await fetch("http://localhost:8081/user/signIn", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: 'include', // Allows cookies to be sent
            body: JSON.stringify({ username, password }),
        });

        if (response.ok) {
            // JWT token will be automatically set by the backend cookie, no need to handle it manually
            console.log("jwt", Cookies.get("jwt"));
            router.push("/"); // Redirect to the main page after successful login
        } else {
            const responseMessage = await response.json();
            setErrorMessage(responseMessage.error || "Authentication failed");
            setTimeout(() => setErrorMessage(null), 7000);
        }
    };

    return (
        <>
            {errorMessage && (
                <Notification
                    color="red"
                    icon={<IconX size="1.1rem" />}
                    style={{
                        position: 'fixed',
                        top: '20px',
                        right: '20px',
                        zIndex: 999,
                    }}
                    onClose={() => setErrorMessage(null)}
                >
                    {errorMessage}
                </Notification>
            )}
            <Container size="sm" mt={50}>
                <Card shadow="sm" p="lg" radius="md" withBorder>
                    <Title ta="center" order={2} mb="lg">Login</Title>
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
                    <Button variant="filled" color="blue" size="xl" radius="xl" onClick={handleSignIn}>
                        Login
                    </Button>
                    <Space h="md" />
                    <Group justify="center" mt="md">
                        <Text size="sm">Do not have an account?</Text>
                        <Link href="/sign-up" passHref>
                            <Text size="sm" td="underline"
                                  style={{ cursor: "pointer", fontWeight: "bold", color: "black" }}>
                                Create an account
                            </Text>
                        </Link>
                    </Group>
                </Card>
            </Container>
        </>
    );
}
