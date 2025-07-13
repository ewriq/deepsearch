
import { writable } from 'svelte/store';
import axios from 'axios';
import { marked } from 'marked';


export const searchTerm = writable('');
export const searchResult = writable(null);
export const errorMessage = writable(null);
export const isLoading = writable(false);

export async function performSearch() {
  errorMessage.set(null);
  searchResult.set(null);
  isLoading.set(true);

  let term = '';
  const unsubscribe = searchTerm.subscribe(value => {
    term = value;
  });
  unsubscribe();

  if (term.trim() === '') {
    errorMessage.set('Lütfen bir arama terimi girin.');
    isLoading.set(false);
    return;
  }

  try {
    const response = await axios.get(`http://localhost:3000/search/${encodeURIComponent(term)}`);
    const data = response.data;

    if (data.status === 'success') {
      searchResult.set(data.data || 'Sonuç bulunamadı.');
    } else if (data.status === 'no content') {
      errorMessage.set(data.message || 'Arama terimi eksik.');
    } else {
      errorMessage.set(data.message || 'Arama sırasında bir hata oluştu.');
    }
  } catch (error) {
    if (axios.isAxiosError && axios.isAxiosError(error)) { 
      if (error.response) {
        errorMessage.set(error.response.data?.message || `API Hatası: ${error.response.status}`);
      } else if (error.request) {
        errorMessage.set('Sunucuya ulaşılamıyor. Ağ bağlantınızı kontrol edin.');
      } else {
        errorMessage.set(error.message || 'Bilinmeyen bir Axios hatası oluştu.');
      }
    } else if (error instanceof Error) {
      errorMessage.set(error.message);
    } else {
      errorMessage.set('Bilinmeyen bir hata oluştu.');
    }
    console.error('API hatası:', error);
  } finally {
    isLoading.set(false);
  }
}


export function handleKeyPress(event) {
  if (event.key === 'Enter') {
    performSearch();
  }
}

export function renderMarkdown(text) {
    if (text === null || text === undefined) { // Null veya undefined kontrolü ekle
        return '';
    }
  return marked(text);
}