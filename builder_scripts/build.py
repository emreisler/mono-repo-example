import os


def get_files() -> []:
    return os.listdir()


def filter_folders(list: []) -> []:
    return [list for element in list if not "." not in element]


def create_mod_files(folder_list: []):
    """
    creates go mod files in modules
    :param folder_list:
    :return:
    """


def main():
    directories=[d for d in os.listdir(os.getcwd()) if os.path.isdir(d)]
    create_mod_files(directories)


if __name__ == "__main__":
    main()
