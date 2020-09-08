#include <gtk/gtk.h>
#include <libcomasync.h>

#include "status.h"

static void activate(GtkApplication* app, gpointer user_data) {
    GtkWidget *window = gtk_application_window_new(app);
    gtk_window_set_title(GTK_WINDOW(window), "Com Async Library");
    gtk_window_set_position(GTK_WINDOW(window), GTK_WIN_POS_CENTER);
    gtk_container_set_border_width(GTK_CONTAINER(window), 10);

    GtkWidget* grid = gtk_grid_new();
    gtk_container_add(GTK_CONTAINER(window), grid);

    GtkWidget* input_grid = gtk_grid_new();
    GtkWidget* input_label = gtk_label_new_with_mnemonic("Input text:");
    gtk_label_set_xalign(GTK_LABEL(input_label), 0.0);
    GtkWidget* input = gtk_entry_new();
    gtk_widget_set_margin_end(GTK_WIDGET(input), 20);
    gtk_entry_set_placeholder_text(GTK_ENTRY(input), "Input...");
    gtk_grid_attach(GTK_GRID(input_grid), input_label, 0, 0, 1, 1);
    gtk_grid_attach(GTK_GRID(input_grid), input, 0, 1, 1, 1);

    GtkWidget* output_grid = gtk_grid_new();
    GtkWidget* output_label = gtk_label_new_with_mnemonic("Output text:");
    gtk_label_set_xalign(GTK_LABEL(output_label), 0.0);
    GtkWidget* output = gtk_entry_new();
    gtk_entry_set_placeholder_text(GTK_ENTRY(output), "Output...");
    gtk_grid_attach(GTK_GRID(output_grid), output_label, 0, 0, 1, 1);
    gtk_grid_attach(GTK_GRID(output_grid), output, 0, 1, 1, 1);

    GtkWidget* tree = create_status_view();
    gtk_widget_set_margin_top(GTK_WIDGET(tree), 20);

    gtk_grid_attach(GTK_GRID(grid), input_grid, 0, 0, 1, 1);
    gtk_grid_attach(GTK_GRID(grid), output_grid, 1, 0, 1, 1);
    gtk_grid_attach(GTK_GRID(grid), tree, 0, 1, 2, 1);

    gtk_widget_show_all(window);
}

int main(int argc, char **argv) {
    GtkApplication* app = gtk_application_new("org.gtk.comasync", G_APPLICATION_FLAGS_NONE);
    g_signal_connect(app, "activate", G_CALLBACK(activate), NULL);

    int status = g_application_run(G_APPLICATION(app), argc, argv);
    g_object_unref(app);
    return status;
}