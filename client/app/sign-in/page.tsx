"use client";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Button, Card, TextInput, Text, Container, Title, Group, Space } from "@mantine/core";
import '@mantine/core/styles.css';
import Link from "next/link";
import Cookies from "js-cookie";
import ErrorNotification from "../components/ErrorNotification";

export default function SignInPage() {
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

    const handleSignIn = async () => {
        const response = await fetch("http://localhost:8081/user/signIn", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: 'include',
            body: JSON.stringify({ username, password }),
        });

        if (response.ok) {
            const data = await response.json();
            const expirationTimestampInSeconds = data.refresh;
            Cookies.set("refresh", "valid", { expires: new Date(expirationTimestampInSeconds * 1000) });
            router.push("/");
        } else {
            const responseMessage = await response.json();
            setErrorMessage(responseMessage.error || "Authentication failed");
            setTimeout(() => setErrorMessage(null), 7000);
        }
    };

    return (
        <>
            <ErrorNotification errorMessage={errorMessage} onClose={() => setErrorMessage(null)} />

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
