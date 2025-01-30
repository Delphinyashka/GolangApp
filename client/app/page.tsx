"use client";
import {useState, useEffect} from "react";
import {useRouter} from "next/navigation";
import {Button, Card, Text, Container, Title, Space, Table, Pagination, Notification} from "@mantine/core";
import '@mantine/core/styles.css';
import Cookies from "js-cookie";
import {IconX} from "@tabler/icons-react";

export default function MainPage() {
    const [orders, setOrders] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);
    const [totalOrders, setTotalOrders] = useState(0);
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const router = useRouter();

    useEffect(() => {
        const jwtToken = Cookies.get("jwt");
        const refreshToken = Cookies.get("refresh");
        if (jwtToken) {
            fetchOrders(jwtToken, currentPage);
        } else if (!jwtToken && refreshToken) {
            handleRefresh();
        } else {
            router.push("/sign-in");
        }
    }, [currentPage, router]);

    const handleRefresh = async () => {
        const response = await fetch("http://localhost:8081/user/refresh", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            credentials: 'include',
        });
    };

    const fetchOrders = async (token: string, page: number) => {
        const response = await fetch(`http://localhost:8082/api/orders?page=${page}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
            credentials: 'include',
        });

        if (response.ok) {
            const data = await response.json();
            setOrders(data.orders);
            setTotalOrders(data.total);
        } else {
            const responseMessage = await response.json();
            setErrorMessage(responseMessage.error || "Authentication failed");
            setTimeout(() => setErrorMessage(null), 7000);
        }
    };

    const handleSignOut = async () => {
        Cookies.remove("refresh");
        Cookies.remove("jwt");
        router.push("/sign-in");
    };

    return (
        <>
            {errorMessage && (
                <Notification
                    color="red"
                    icon={<IconX size="1.1rem"/>}
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

                <Card shadow="sm" p="lg" radius="md" withBorder mt="xl">
                    <Title order={3}>Your Orders</Title>
                    <Table mt="md">
                        <thead>
                        <tr>
                            <th>Product Name</th>
                            <th>Client Name</th>
                            <th>Price</th>
                            <th>Order ID</th>
                        </tr>
                        </thead>
                        <tbody>
                        {orders.map((order: any) => (
                            <tr key={order.id}>
                                <td>{order.productName}</td>
                                <td>{order.clientName}</td>
                                <td>{order.price}</td>
                                <td>{order.id}</td>
                            </tr>
                        ))}
                        </tbody>
                    </Table>
                    <Pagination
                        page={currentPage}
                        onChange={setCurrentPage}
                        total={Math.ceil(totalOrders / 10)}
                        mt="md"
                    />
                </Card>

                <Card shadow="sm" p="lg" radius="md" withBorder>
                    <Button variant="outline" color="red" size="xl" radius="xl" onClick={handleSignOut}>
                        Logout
                    </Button>
                </Card>
            </Container>
        </>
    );
}
