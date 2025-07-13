<script lang="js">
    import {
      searchTerm,
      searchResult,
      errorMessage,
      isLoading,
      performSearch,
      handleKeyPress,
      renderMarkdown
    } from '../lib/index'; 
</script>

<div class="flex flex-col items-center justify-center min-h-screen p-6 bg-gray-900 text-gray-100 font-sans">
  <h1 class="text-5xl font-extrabold text-white mb-6">DeepSearch</h1>
  <p class="text-lg text-gray-300 mb-8">Arama yapmak için bir query girin:</p>

  <div class="flex gap-4 mb-10 w-full max-w-xl">
    <input
      type="text"
      placeholder="Örn: ewriq etc."
      bind:value={$searchTerm} 
      on:keypress={handleKeyPress}
      disabled={$isLoading}
      class="flex-grow max-w-md p-3 rounded-lg border-2 border-gray-600 bg-gray-700 text-gray-100 text-lg
             focus:border-blue-500 focus:ring-2 focus:ring-blue-500 outline-none transition-all duration-300"
    />
    <button
      on:click={performSearch}
      disabled={$isLoading}
      class="bg-purple-700 text-white font-bold py-3 px-6 rounded-lg shadow-lg
             hover:bg-purple-600 active:bg-purple-800 focus:outline-none focus:ring-2 focus:ring-purple-500
             transition-all duration-300 disabled:bg-gray-600 disabled:cursor-not-allowed flex items-center justify-center gap-2"
    >
      {#if $isLoading}
        <svg class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        Aranıyor...
      {:else}
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="11" cy="11" r="8"></circle>
          <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
        </svg>
        Ara
      {/if}
    </button>
  </div>

  {#if $errorMessage}
    <p class="bg-red-700 text-white font-bold p-4 rounded-lg shadow-md mt-6 w-full max-w-xl text-center">{$errorMessage}</p>
  {:else if $searchResult}
    <div class="bg-gray-800 border border-gray-700 rounded-xl p-8 mt-8 w-full max-w-3xl shadow-xl">
      <h2 class="text-3xl font-semibold text-white mb-6 border-b border-gray-700 pb-4">Arama Sonucu:</h2>
      <div class="rendered-markdown bg-gray-700 p-6 rounded-lg overflow-x-auto text-gray-200 leading-relaxed text-left">
        {@html renderMarkdown($searchResult)}
      </div>
    </div>
  {:else if !$isLoading && $searchTerm.trim() !== ''}
    <p class="text-gray-400 italic mt-6 text-lg">Sonuç bulunamadı veya henüz arama yapılmadı.</p>
  {/if}
</div>

<style>
  .rendered-markdown p {
      margin-bottom: 1em;
  }
  .rendered-markdown h1, .rendered-markdown h2, .rendered-markdown h3,
  .rendered-markdown h4, .rendered-markdown h5, .rendered-markdown h6 {
      color: #e0e6ed; 
      margin-top: 1.5em;
      margin-bottom: 0.8em;
      font-weight: bold;
  }
  .rendered-markdown h1 { font-size: 2em; }
  .rendered-markdown h2 { font-size: 1.75em; }
  .rendered-markdown h3 { font-size: 1.5em; }

  .rendered-markdown ul, .rendered-markdown ol {
      list-style-position: inside;
      padding-left: 1.5em;
      margin-bottom: 1em;
  }
  .rendered-markdown li {
      margin-bottom: 0.5em;
  }
  .rendered-markdown pre {
      background-color: #1a202c;
      color: #a0aec0;
      padding: 1em;
      border-radius: 0.5rem;
      overflow-x: auto;
      margin-bottom: 1em;
  }
  .rendered-markdown code {
      background-color: #2d3748; 
      color: #f7fafc;
      padding: 0.2em 0.4em;
      border-radius: 0.25rem;
      font-family: 'Fira Code', 'JetBrains Mono', monospace; 
      font-size: 0.9em;
  }
  .rendered-markdown a {
      color: #63b3ed; 
      text-decoration: underline;
  }
</style>