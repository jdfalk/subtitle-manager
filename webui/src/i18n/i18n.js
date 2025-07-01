// file: webui/src/i18n/i18n.js
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174003

import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';

// Translation resources
const resources = {
  en: {
    translation: {
      // Navigation
      'nav.dashboard': 'Dashboard',
      'nav.library': 'Media Library',
      'nav.convert': 'Convert',
      'nav.translate': 'Translate',
      'nav.extract': 'Extract',
      'nav.history': 'History',
      'nav.wanted': 'Wanted',
      'nav.settings': 'Settings',
      'nav.system': 'System',
      'nav.scheduling': 'Scheduling',
      'nav.tags': 'Tag Management',
      'nav.users': 'User Management',
      
      // Common
      'common.save': 'Save',
      'common.cancel': 'Cancel',
      'common.delete': 'Delete',
      'common.edit': 'Edit',
      'common.close': 'Close',
      'common.refresh': 'Refresh',
      'common.loading': 'Loading...',
      'common.error': 'Error',
      'common.success': 'Success',
      'common.warning': 'Warning',
      'common.info': 'Information',
      
      // Settings
      'settings.title': 'Settings',
      'settings.language': 'Language',
      'settings.theme': 'Theme',
      'settings.theme.light': 'Light',
      'settings.theme.dark': 'Dark',
      'settings.theme.auto': 'Auto',
      
      // App Title
      'app.title': 'Subtitle Manager',
      'app.subtitle': 'Manage your subtitles',
      
      // Dashboard
      'dashboard.title': 'Dashboard',
      'dashboard.overview': 'Overview',
      'dashboard.recent': 'Recent Activity',
      
      // Media Library
      'library.title': 'Media Library',
      'library.scan': 'Scan Library',
      'library.filter': 'Filter',
      'library.search': 'Search media...',
      
      // Convert
      'convert.title': 'Convert Subtitles',
      'convert.input': 'Input File',
      'convert.output': 'Output File',
      'convert.format': 'Output Format',
      'convert.start': 'Start Conversion',
      
      // Translate
      'translate.title': 'Translate Subtitles',
      'translate.source': 'Source Language',
      'translate.target': 'Target Language',
      'translate.service': 'Translation Service',
      'translate.start': 'Start Translation',
      
      // Extract
      'extract.title': 'Extract Subtitles',
      'extract.media': 'Media File',
      'extract.tracks': 'Subtitle Tracks',
      'extract.start': 'Start Extraction',
      
      // History
      'history.title': 'History',
      'history.date': 'Date',
      'history.operation': 'Operation',
      'history.file': 'File',
      'history.status': 'Status',
      
      // Wanted
      'wanted.title': 'Wanted Subtitles',
      'wanted.movie': 'Movie',
      'wanted.episode': 'Episode',
      'wanted.language': 'Language',
      'wanted.search': 'Search',
      
      // System
      'system.title': 'System Information',
      'system.version': 'Version',
      'system.uptime': 'Uptime',
      'system.memory': 'Memory Usage',
      'system.logs': 'Logs',
    }
  },
  es: {
    translation: {
      // Navigation
      'nav.dashboard': 'Panel',
      'nav.library': 'Biblioteca de Medios',
      'nav.convert': 'Convertir',
      'nav.translate': 'Traducir',
      'nav.extract': 'Extraer',
      'nav.history': 'Historial',
      'nav.wanted': 'Buscados',
      'nav.settings': 'Configuración',
      'nav.system': 'Sistema',
      'nav.scheduling': 'Programación',
      'nav.tags': 'Gestión de Etiquetas',
      'nav.users': 'Gestión de Usuarios',
      
      // Common
      'common.save': 'Guardar',
      'common.cancel': 'Cancelar',
      'common.delete': 'Eliminar',
      'common.edit': 'Editar',
      'common.close': 'Cerrar',
      'common.refresh': 'Actualizar',
      'common.loading': 'Cargando...',
      'common.error': 'Error',
      'common.success': 'Éxito',
      'common.warning': 'Advertencia',
      'common.info': 'Información',
      
      // Settings
      'settings.title': 'Configuración',
      'settings.language': 'Idioma',
      'settings.theme': 'Tema',
      'settings.theme.light': 'Claro',
      'settings.theme.dark': 'Oscuro',
      'settings.theme.auto': 'Automático',
      
      // App Title
      'app.title': 'Gestor de Subtítulos',
      'app.subtitle': 'Gestiona tus subtítulos',
      
      // Dashboard
      'dashboard.title': 'Panel',
      'dashboard.overview': 'Resumen',
      'dashboard.recent': 'Actividad Reciente',
      
      // Media Library
      'library.title': 'Biblioteca de Medios',
      'library.scan': 'Escanear Biblioteca',
      'library.filter': 'Filtrar',
      'library.search': 'Buscar medios...',
      
      // Convert
      'convert.title': 'Convertir Subtítulos',
      'convert.input': 'Archivo de Entrada',
      'convert.output': 'Archivo de Salida',
      'convert.format': 'Formato de Salida',
      'convert.start': 'Iniciar Conversión',
      
      // Translate
      'translate.title': 'Traducir Subtítulos',
      'translate.source': 'Idioma de Origen',
      'translate.target': 'Idioma de Destino',
      'translate.service': 'Servicio de Traducción',
      'translate.start': 'Iniciar Traducción',
      
      // Extract
      'extract.title': 'Extraer Subtítulos',
      'extract.media': 'Archivo de Medios',
      'extract.tracks': 'Pistas de Subtítulos',
      'extract.start': 'Iniciar Extracción',
      
      // History
      'history.title': 'Historial',
      'history.date': 'Fecha',
      'history.operation': 'Operación',
      'history.file': 'Archivo',
      'history.status': 'Estado',
      
      // Wanted
      'wanted.title': 'Subtítulos Buscados',
      'wanted.movie': 'Película',
      'wanted.episode': 'Episodio',
      'wanted.language': 'Idioma',
      'wanted.search': 'Buscar',
      
      // System
      'system.title': 'Información del Sistema',
      'system.version': 'Versión',
      'system.uptime': 'Tiempo de Actividad',
      'system.memory': 'Uso de Memoria',
      'system.logs': 'Registros',
    }
  },
  fr: {
    translation: {
      // Navigation
      'nav.dashboard': 'Tableau de bord',
      'nav.library': 'Bibliothèque multimédia',
      'nav.convert': 'Convertir',
      'nav.translate': 'Traduire',
      'nav.extract': 'Extraire',
      'nav.history': 'Historique',
      'nav.wanted': 'Recherchés',
      'nav.settings': 'Paramètres',
      'nav.system': 'Système',
      'nav.scheduling': 'Planification',
      'nav.tags': 'Gestion des étiquettes',
      'nav.users': 'Gestion des utilisateurs',
      
      // Common
      'common.save': 'Enregistrer',
      'common.cancel': 'Annuler',
      'common.delete': 'Supprimer',
      'common.edit': 'Modifier',
      'common.close': 'Fermer',
      'common.refresh': 'Actualiser',
      'common.loading': 'Chargement...',
      'common.error': 'Erreur',
      'common.success': 'Succès',
      'common.warning': 'Avertissement',
      'common.info': 'Information',
      
      // Settings
      'settings.title': 'Paramètres',
      'settings.language': 'Langue',
      'settings.theme': 'Thème',
      'settings.theme.light': 'Clair',
      'settings.theme.dark': 'Sombre',
      'settings.theme.auto': 'Automatique',
      
      // App Title
      'app.title': 'Gestionnaire de Sous-titres',
      'app.subtitle': 'Gérez vos sous-titres',
      
      // Dashboard
      'dashboard.title': 'Tableau de bord',
      'dashboard.overview': 'Aperçu',
      'dashboard.recent': 'Activité récente',
      
      // Media Library
      'library.title': 'Bibliothèque multimédia',
      'library.scan': 'Scanner la bibliothèque',
      'library.filter': 'Filtrer',
      'library.search': 'Rechercher des médias...',
      
      // Convert
      'convert.title': 'Convertir les sous-titres',
      'convert.input': 'Fichier d\'entrée',
      'convert.output': 'Fichier de sortie',
      'convert.format': 'Format de sortie',
      'convert.start': 'Démarrer la conversion',
      
      // Translate
      'translate.title': 'Traduire les sous-titres',
      'translate.source': 'Langue source',
      'translate.target': 'Langue cible',
      'translate.service': 'Service de traduction',
      'translate.start': 'Démarrer la traduction',
      
      // Extract
      'extract.title': 'Extraire les sous-titres',
      'extract.media': 'Fichier multimédia',
      'extract.tracks': 'Pistes de sous-titres',
      'extract.start': 'Démarrer l\'extraction',
      
      // History
      'history.title': 'Historique',
      'history.date': 'Date',
      'history.operation': 'Opération',
      'history.file': 'Fichier',
      'history.status': 'Statut',
      
      // Wanted
      'wanted.title': 'Sous-titres recherchés',
      'wanted.movie': 'Film',
      'wanted.episode': 'Épisode',
      'wanted.language': 'Langue',
      'wanted.search': 'Rechercher',
      
      // System
      'system.title': 'Informations système',
      'system.version': 'Version',
      'system.uptime': 'Temps de fonctionnement',
      'system.memory': 'Utilisation de la mémoire',
      'system.logs': 'Journaux',
    }
  }
};

i18n
  .use(initReactI18next)
  .init({
    resources,
    lng: 'en', // default language
    fallbackLng: 'en',
    interpolation: {
      escapeValue: false, // React already does escaping
    },
  });

export default i18n;