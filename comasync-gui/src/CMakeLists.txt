find_package(PkgConfig REQUIRED)
pkg_check_modules(GTK REQUIRED gtk+-3.0 gmodule-2.0)

add_executable(comasync-gui main.c status.h)
add_subdirectory(libcomasync)
add_definitions(${GTK_CFLAGS_OTHER})

target_include_directories(comasync-gui PUBLIC ${GTK_INCLUDE_DIRS})
target_link_directories(comasync-gui PUBLIC ${GTK_LIBRARY_DIRS})
target_link_libraries(comasync-gui PUBLIC libcomasync ${GTK_LIBRARIES})
