const input = document.querySelector("input");

input.addEventListener("keydown", (event) => {
  if (event.key === "Enter") {
    const value = input.value.trim();
    if (value) {
      const newUrl = window.location.origin + "/" + value;
      window.location.href = newUrl;
      // alert(value);
    }
  }
});
