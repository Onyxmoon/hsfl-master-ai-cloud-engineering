import { handleErrors } from '../../../assets/helper/handleErrors';
import { isAuthenticated } from "../../../store";

interface Merchant {
    id: number;
    name: string;
    role?: number;
    productsCount?: number;
}

interface Price {
    userId: number;
    productId: number;
    price: number;
}

interface Product {
    id: number,
    description: string,
    ean: number,
}

export const load = async (context: { params: { id: string } }): Promise<Promise<object> | undefined> => {
    if (! isAuthenticated) {
        return;
    }

    const { id } = context.params;
    const apiUrlMerchant: string = `/api/v1/user/${id}`;
    const apiUrlPrices: string = `/api/v1/price/user/${id}`;

    try {
        const [merchant, prices] = await Promise.all([
            fetch(apiUrlMerchant).then(handleErrors) as Promise<Merchant>,
            fetch(apiUrlPrices).then(handleErrors) as Promise<Price[]>,
        ]);

        let sortedProducts: Product[] = [];

        if (prices) {
            const uniqueProductIds: number[] = Array.from(new Set(prices.map(price => price.productId)));
            const productsPromises: Promise<Product>[] = uniqueProductIds.map(productId =>
                fetch(`/api/v1/product/${productId}`).then(handleErrors) as Promise<Product>
            );

            const products: Product[] = await Promise.all(productsPromises);
            sortedProducts = products.sort(
                (a: Product, b: Product) => a.description.localeCompare(b.description));
        }

        return {
            merchant: merchant,
            prices: prices ?? [],
            products: sortedProducts,
            metaTitle: merchant?.name
        };
    } catch (error) {
        return {
            merchant: null,
            prices: [],
            products: [],
            metaTitle: 'Leider ist ein Fehler aufgetreten.',
        };
    }
};

