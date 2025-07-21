import { get, writable } from 'svelte/store';
import { marked } from 'marked';
export const searchTerm = writable('');
export const searchResult = writable(null);
export const errorMessage = writable(null);
export const isLoading = writable(false);
export const progressPercent = writable(0);
export const progressMessage = writable('');

let socket; 

export function performSearch() {
  isLoading.set(true);
  errorMessage.set(null);
  searchResult.set(null);
  progressPercent.set(0);
  progressMessage.set('Bağlantı kuruluyor...');


  let term = '';
  const unsubscribe = searchTerm.subscribe(value => { term = value; });
  unsubscribe();

  if (term.trim() === '') {
    errorMessage.set('Lütfen bir arama terimi girin.');
    isLoading.set(false);
    return;
  }
  socket = new WebSocket(`ws://localhost:3000/ws/search`);
  socket.onopen = () => {
    console.log('WebSocket bağlantısı başarıyla kuruldu.');
    progressMessage.set('Bağlantı başarılı. Sorgu gönderiliyor...');
    socket.send(term);
  };


  socket.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    console.log('Sunucudan mesaj:', msg);
    switch (msg.status) {
      case 'progress':
        progressPercent.set(msg.progress);
        progressMessage.set(msg.message);
        break;

      case 'complete':
        progressPercent.set(100);
        progressMessage.set('İşlem tamamlandı!');
        searchResult.set(msg.data);
        isLoading.set(false);
        socket.close();
        break;

      case 'error':
        errorMessage.set(msg.message);
        progressPercent.set(100); 
        isLoading.set(false);
        socket.close(); 
        break;
    }
  };

  socket.onerror = (error) => {
    console.error('WebSocket hatası:', error);
    errorMessage.set('Sunucuya bağlanılamadı. Lütfen sunucunun çalıştığından emin olun.');
    isLoading.set(false);
  };

  socket.onclose = () => {
    console.log('WebSocket bağlantısı kapandı.');
    if (get()) {
      isLoading.set(false);
    }
  };
}


export function handleKeyPress(event) {
  if (event.key === 'Enter' && !get()) {
    performSearch();
  }
}

export function renderMarkdown(text) {
  if (text === null || text === undefined) {
    return '';
  }
  return marked(text);
}