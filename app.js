const API_BASE_URL = "http://localhost:8080/api/cr/v1";

// Utility function to show alert messages
function showAlert(message, type) {
  const alertPlaceholder = document.getElementById("alertPlaceholder");
  alertPlaceholder.innerHTML = `
    <div class="alert alert-${type} alert-dismissible fade show" role="alert">
      ${message}
      <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
    </div>
  `;
}

// Add Item
document.getElementById("addItemForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  const itemName = document.getElementById("itemName").value;
  const description = document.getElementById("itemDescription").value;

  try {
    const response = await fetch(`${API_BASE_URL}/additem`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ itemName, description }),
    });
    const data = await response.json();
    if (response.ok) {
      showAlert(data.message, "success");
      e.target.reset();
    } else {
      showAlert(data.message || "Failed to add item.", "danger");
    }
  } catch (error) {
    console.error("Error adding item:", error);
    showAlert("An error occurred. Please try again.", "danger");
  }
});

// Refresh Items
document.getElementById("refreshItemsList").addEventListener("click", async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/getitems`);
    const data = await response.json();
    const itemsList = document.getElementById("itemsList");
    itemsList.innerHTML = "";
    data.items.forEach((item) => {
      const listItem = document.createElement("li");
      listItem.classList.add("list-group-item");
      listItem.textContent = `ID: ${item.itemId}, Name: ${item.itemName}, Description: ${item.description}`;
      itemsList.appendChild(listItem);
    });
  } catch (error) {
    console.error("Error fetching items:", error);
    showAlert("Failed to fetch items.", "danger");
  }
});

// Add Client
document.getElementById("addClientForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  const clientName = document.getElementById("clientName").value;
  const phone = document.getElementById("clientPhone").value;
  const email = document.getElementById("clientEmail").value;

  try {
    const response = await fetch(`${API_BASE_URL}/addclient`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ clientName, phone, email }),
    });
    const data = await response.json();
    if (response.ok) {
      showAlert(data.message, "success");
      e.target.reset();
    } else {
      showAlert(data.message || "Failed to add client.", "danger");
    }
  } catch (error) {
    console.error("Error adding client:", error);
    showAlert("An error occurred. Please try again.", "danger");
  }
});

// Refresh Clients
document.getElementById("refreshClientsList").addEventListener("click", async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/getclients`);
    const data = await response.json();
    const clientsList = document.getElementById("clientsList");
    clientsList.innerHTML = "";
    data.clients.forEach((client) => {
      const listItem = document.createElement("li");
      listItem.classList.add("list-group-item");
      listItem.textContent = `ID: ${client.clientId}, Name: ${client.clientName}, Phone: ${client.phone}, Email: ${client.email}`;
      clientsList.appendChild(listItem);
    });
  } catch (error) {
    console.error("Error fetching clients:", error);
    showAlert("Failed to fetch clients.", "danger");
  }
});

// Add Transaction
document.getElementById("addTransactionForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  const clientId = document.getElementById("transactionClientId").value;
  const itemId = document.getElementById("transactionItemId").value;

  try {
    const response = await fetch(`${API_BASE_URL}/addtransaction`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ clientId: parseInt(clientId), itemId: parseInt(itemId) }),
    });
    const data = await response.json();
    if (response.ok) {
      showAlert(data.message, "success");
      e.target.reset();
    } else {
      showAlert(data.message || "Failed to add transaction.", "danger");
    }
  } catch (error) {
    console.error("Error adding transaction:", error);
    showAlert("An error occurred. Please try again.", "danger");
  }
});

// Refresh Transactions
document.getElementById("refreshTransactionsList").addEventListener("click", async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/gettransactions`);
    const data = await response.json();
    const transactionsList = document.getElementById("transactionsList");
    transactionsList.innerHTML = "";
    data.transactions.forEach((transaction) => {
      const listItem = document.createElement("li");
      listItem.classList.add("list-group-item");
      listItem.textContent = `Transaction ID: ${transaction.transactionId}, Client ID: ${transaction.clientId}, Item ID: ${transaction.itemId}`;
      transactionsList.appendChild(listItem);
    });
  } catch (error) {
    console.error("Error fetching transactions:", error);
    showAlert("Failed to fetch transactions.", "danger");
  }
});

// Get Transaction Details
document.getElementById("fetchTransactionDetails").addEventListener("click", async () => {
  const transactionId = document.getElementById("transactionId").value;
  if (!transactionId) {
    showAlert("Please enter a transaction ID.", "warning");
    return;
  }

  try {
    const response = await fetch(`${API_BASE_URL}/gettransactiondetails`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ transactionId: parseInt(transactionId) }),
    });
    const data = await response.json();
    if (response.ok) {
      const details = `
        Transaction ID: ${data.transactionId}
        Client ID: ${data.clientId}, Name: ${data.clientName}, Phone: ${data.phone}, Email: ${data.email}
        Item ID: ${data.itemId}, Name: ${data.itemName}, Description: ${data.itemDescription}
      `;
      showAlert(details.replace(/\n/g, "<br>"), "info");
    }else {
      showAlert(data.message || "Failed to fetch transaction details.", "danger");
    }
  } catch (error) {
    console.error("Error fetching transaction details:", error);
    showAlert("An error occurred. Please try again.", "danger");
  }
});
