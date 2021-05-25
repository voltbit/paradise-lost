from argparse import ArgumentParser
from urllib.error import HTTPError
import urllib.request
import gzip
import sys

class URLs():
    STABLE = 'http://ftp.uk.debian.org/debian/dists/stable/main'

class Debstat():
    """Utility object used to download information about Debian packages based on the architecture
    they were built for.
    """

    @classmethod
    def retrieve_data(cls, arch: str):
        """Connects to the webserver and downloads the file based on the provided argument or throws
        an error if the architecture name is not found/connection could not be established.
        """
        try:
            file_name, _ = urllib.request.urlretrieve(f'{URLs.STABLE}/Contents-{arch}.gz')
        except HTTPError as err:
            if err.code > 399 and err.code < 500:
                print('Resource not found. Make sure the arch identifier is correct.')
                sys.exit(1)
            else:
                print('Could not connect to the web server.')
                sys.exit(1)
        return file_name

    @classmethod
    def get_map(cls, raw_data_path: str):
        """Reads the raw data and generates a short report"""
        data = {}
        with gzip.open(raw_data_path) as f_handle:
            pkg_data = f_handle.readline()
            while pkg_data:
                pkg_data = pkg_data.decode('UTF-8')
                if pkg_data.find('/') != -1:
                    pkg_data = pkg_data.split()
                    if len(pkg_data) > 1:
                        if len(pkg_data) > 2:
                            pkg_data = [" ".join(pkg_data[:-1]), pkg_data[-1]]
                        for package in pkg_data[1].split(','):
                            if package in data:
                                data[package].append(pkg_data[0])
                            else:
                                data[package] = []
                                data[package].append(pkg_data[0])
                    else:
                        print(f'Outlier: {pkg_data}')
                pkg_data = f_handle.readline()
        return data

    @classmethod
    def get_top_package_by_file_count(cls, arch, n=10):
        raw_data = cls.retrieve_data(arch)
        data_map = cls.get_map(raw_data)
        top_n = sorted([(x, len(y)) for x, y in data_map.items()], key=lambda x: x[1], reverse=True)
        count = 1
        for entry in top_n[:n]:
            print(f'{count}. {entry[0]}\t\t{entry[1]}')
            count += 1
        return top_n


def main():
    parser = ArgumentParser(description='Statistics about packages for an architecture')
    parser.add_argument('arch', metavar='architecture',
                        help='string representing the desired type of architecture')
    args = parser.parse_args()
    Debstat.get_top_package_by_file_count(args.arch)

if __name__ == '__main__':
    main()
