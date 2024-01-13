import { handleErrors } from "./handleErrors";

interface Product {
    id: number,
    description: string,
    ean: number,
}

export async function sortProducts(items: any): Promise<Product[]> {
    if (!items) return [];

    const uniqueProductIds: number[] = Array.from(new Set(items.map((item: { productId: number }) => item.productId)));
    const productsPromises: Promise<Product | null>[] = uniqueProductIds.map(productId =>
        fetch(`/api/v1/product/${productId}`)
            .then(handleErrors)
            .then(product => (product instanceof Error ? null : product)) as Promise<Product | null>
    );

    const products: Product[] = (await Promise.all(productsPromises)).filter((product) => product !== null) as Product[];

    return products.sort(
        (a: Product, b: Product) => a.description.localeCompare(b.description)
    );
}
