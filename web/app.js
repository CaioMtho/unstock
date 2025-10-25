const API_URL = "http://localhost:8080/api/products";

async function loadProducts() {
  const res = await fetch(API_URL + "/");
  const products = await res.json();
  renderProducts(products);
}

function renderProducts(products) {
  const tbody = document.querySelector("#products-table tbody");
  tbody.innerHTML = "";
  products.forEach(p => {
    const tr = document.createElement("tr");
    if (p.stock < p.min_stock) tr.classList.add("low-stock");
    tr.innerHTML = `
      <td>${p.id}</td>
      <td>${p.name}</td>
      <td>${p.price}</td>
      <td>${p.stock}</td>
      <td>${p.min_stock}</td>
      <td>
        <button onclick="updateStock(${p.id}, 1)">+1</button>
        <button onclick="updateStock(${p.id}, -1)">-1</button>
        <button onclick="deleteProduct(${p.id})">Remover</button>
      </td>
    `;
    tbody.appendChild(tr);
  });
}

document.getElementById("add-product-form").addEventListener("submit", async e => {
  e.preventDefault();
  const formData = new FormData(e.target);
  const body = Object.fromEntries(formData.entries());
  body.price = parseFloat(body.price);
  body.stock = parseInt(body.stock);
  body.min_stock = parseInt(body.min_stock);

  await fetch(API_URL + "/", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body)
  });
  e.target.reset();
  loadProducts();
});

async function updateStock(id, value) {
  await fetch(`${API_URL}/${id}/stock`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ value })
  });
  loadProducts();
}

async function deleteProduct(id) {
  await fetch(`${API_URL}/${id}`, { method: "DELETE" });
  loadProducts();
}

const socket = new WebSocket("ws://localhost:8080/api/ws/alerts");
socket.onmessage = (event) => {
  const data = JSON.parse(event.data);
  if (data.low_stock) {
    console.log("Produtos com estoque baixo:", data.low_stock);
  }
  if (data.update) {
    console.log("Produto atualizado:", data.update);
    loadProducts();
  }
};

loadProducts();
