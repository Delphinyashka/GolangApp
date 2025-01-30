"use client";

import {useState} from "react";
import {useRouter} from "next/navigation";
import {Button, Card, TextInput, Container, Title, Group, Space, Text, Notification} from "@mantine/core";
import '@mantine/core/styles.css';
import Link from "next/link";
import {z} from "zod";
import {useForm, zodResolver} from "@mantine/form";
import {IconX} from "@tabler/icons-react";

// Zod validation schema (same as in backend)
const schema = z.object({
    username: z.string()
        .min(3, "Username must be at least 3 characters long")
        .max(20, "Username cannot exceed 20 characters")
        .regex(/^[a-zA-Z0-9_]+$/, "Username can only contain letters, numbers, and underscores"),
    password: z.string()
        .min(8, "Password must be at least 8 characters long"),
});

export default function RegisterPage() {
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const router = useRouter();

    const form = useForm({
        validate: zodResolver(schema),
        initialValues: {
            username: "",
            password: "",
        },
    });

    const handleRegister = async (values: { username: string; password: string }) => {
        setErrorMessage(null);

        const response = await fetch("http://localhost:8081/user/signUp", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            credentials: 'include',
            body: JSON.stringify(values),
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
            {errorMessage && (
                <Notification
                    color="red"
                    icon={<IconX size="1.1rem"/>}
                    style={{position: 'fixed', top: '20px', right: '20px', zIndex: 999}}
                    onClose={() => setErrorMessage(null)}
                >
                    {errorMessage}
                </Notification>
            )}
            <Container size="sm" mt={50}>
                <Card shadow="sm" p="lg" radius="md" withBorder>
                    <Title ta="center" order={2} mb="lg">Sign Up</Title>
                    <form onSubmit={form.onSubmit(handleRegister)}>
                        <TextInput
                            label="Username"
                            placeholder="Enter your username"
                            {...form.getInputProps("username")}
                        />
                        <TextInput
                            label="Password"
                            placeholder="Enter your password"
                            type="password"
                            {...form.getInputProps("password")}
                            mt={10}
                        />
                        <Space h="md"/>
                        <Group justify="center">
                            <Button type="submit" variant="filled" color="blue" size="xl" radius="xl">
                                Sign Up
                            </Button>
                        </Group>
                    </form>
                    <Space h="md"/>
                    <Group justify="center" mt="md">
                        <Text size="sm">Already have an account?</Text>
                        <Link href="/sign-in" passHref>
                            <Text size="sm" td="underline"
                                  style={{cursor: "pointer", fontWeight: "bold", color: "black"}}>
                                SIGN IN
                            </Text>
                        </Link>
                    </Group>
                </Card>
            </Container>
        </>
    );
}
