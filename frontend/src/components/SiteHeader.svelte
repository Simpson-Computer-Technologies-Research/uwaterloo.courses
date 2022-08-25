<script>
    export let handleHeaderClick;
</script>

<main>
    <!-- uwaterloo.courses Header -->
    <h1>
        <mark style="color: black; background: none;">uwaterloo.</mark>courses
    </h1>

    <!-- Site Summary -->
    <p style="color: #52525b; margin-left: 1.2%; margin-right: 45%;">
        Try words, phrases, titles, subjects, course numbers, etc. 
        You can also look for exact textual phrases (like "<mark style="color: #6366f1; background: none; cursor: pointer;" on:click={() => handleHeaderClick("engineering")}>engineering</mark>") or filter by specific attributes such as 
        <mark style="color: #6366f1; background: none; cursor: pointer;" on:click={() => handleHeaderClick("@code:cs")}>@code:cs</mark>
    </p>

    <!-- Credits and languages/services used -->
    <p style="color: #52525b; margin-left: 1.2%; margin-right: 45%;">
        Made by 
        <a 
            href="https://github.com/realTristan" 
            style="color: #6366f1;"
            target="_blank"
            rel="noreferrer noopener"
        >
        <nobr>@realTristan</nobr></a> using 

        <!-- Go -->
        <a href="https://go.dev" target="_blank" rel="noopener noreferrer">
            <img src="./images/go_logo.png" alt="" width="44" height="44" style="margin-bottom: -16px; margin-left: -4px;">
        </a>

        and
        
        <!-- Svelte -->
        <a href="https://svelte.dev" style="text-decoration:none;" target="_blank" rel="noopener noreferrer">
            <img src="./images/svelte_logo.png" alt="" width="16.5" height="19.5" style="margin-bottom: -0.6%; margin-left: 4px;"> 
            <mark style="color: #ff3c00; background: none; font-weight: 400; margin-right: 3px;">Svelte</mark>
        </a>

        <!-- Redis 
        <a href="https://redis.io" style="text-decoration:none;" target="_blank" rel="noopener noreferrer">
            <img src="./images/redis_logo.png" alt="" width="18" height="20" style="margin-bottom: -0.6%; margin-left: 5px;">
            <mark style="color: #6b7280; background: none; font-weight: 400;">Redis</mark>
        </a>
        -->
    </p>
</main>

<style>
    h1 {
        color: #8b5cf6;
        font-size: 4em;
        font-weight: 700;
        margin-bottom: 20px;
    }
</style>