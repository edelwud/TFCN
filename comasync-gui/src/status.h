#pragma once

#include <gtk/gtk.h>

enum {
    OPTION_NAME = 0,
    OPTION_VALUE,
    NUM_COLS
};

static GtkTreeModel* migrate_status_store() {
    GtkListStore *store = gtk_list_store_new(NUM_COLS, G_TYPE_STRING, G_TYPE_UINT);
    GtkTreeIter iter;

    gtk_list_store_append(store, &iter);
    gtk_list_store_set(store, &iter, OPTION_NAME, "Heinz El-Mann", OPTION_VALUE, 51, -1);

    gtk_list_store_append(store, &iter);
    gtk_list_store_set(store, &iter, OPTION_NAME, "Jane Doe", OPTION_VALUE, 23, -1);

    gtk_list_store_append(store, &iter);
    gtk_list_store_set(store, &iter, OPTION_NAME, "Joe Bungop", OPTION_VALUE, 91, -1);

    return GTK_TREE_MODEL (store);
}

static GtkWidget* create_status_view() {
    GtkWidget* view = gtk_tree_view_new();

    GtkCellRenderer* renderer = gtk_cell_renderer_text_new();
    gtk_tree_view_insert_column_with_attributes(
            GTK_TREE_VIEW(view), -1, "Option name", renderer, "text", OPTION_NAME, NULL);

    renderer = gtk_cell_renderer_text_new();
    gtk_tree_view_insert_column_with_attributes(
            GTK_TREE_VIEW(view), -1, "Option value", renderer, "text", OPTION_VALUE, NULL);

    GtkTreeModel* model = migrate_status_store();

    gtk_tree_view_set_model(GTK_TREE_VIEW(view), model);

    g_object_unref(model);

    return view;
}

